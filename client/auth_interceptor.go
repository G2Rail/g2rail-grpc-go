package client

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/G2Rail/structs"
	"github.com/Songmu/go-httpdate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var acceptedKinds = []reflect.Kind{reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
	reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String}
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// AuthInterceptor is a client interceptor for authentication
type AuthInterceptor struct {
	authorization string
	apiKey        string
	apiSecret     string
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(apiKey, apiSecret string) (*AuthInterceptor, error) {
	interceptor := &AuthInterceptor{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}

	return interceptor, nil
}

// Unary returns a client interceptor to authenticate unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s", method)

		return invoker(interceptor.attachToken(ctx, req), method, req, reply, cc, opts...)
	}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context, req interface{}) context.Context {
	interceptor.calculateToken(req)
	ctx = metadata.AppendToOutgoingContext(ctx, "from", interceptor.apiKey)
	ctx = metadata.AppendToOutgoingContext(ctx, "date", httpdate.Time2Str(time.Now()))
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.authorization)
}

func (interceptor *AuthInterceptor) calculateToken(req interface{}) {

	fields := structs.Fields(req)
	simpleFields := map[string]interface{}{
		"api_key": interceptor.apiKey,
		"t":       time.Now().Unix(),
	}

	for _, field := range fields {
		if contains(acceptedKinds, field.Kind()) && !strings.HasPrefix(field.Name(), "XXX_") {
			simpleFields[toSnakeCase(field.Name())] = field.Value()
		}
	}

	source := interceptor.buildSignString(&simpleFields)
	interceptor.authorization = getMD5Hash(source)
}

func (interceptor *AuthInterceptor) buildSignString(fields *map[string]interface{}) string {
	var keys []string
	for k := range *fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	source := ""
	for _, k := range keys {
		source += fmt.Sprintf("%s=%v", k, (*fields)[k])
	}
	source += interceptor.apiSecret
	return source
}

func contains(a []reflect.Kind, x reflect.Kind) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
