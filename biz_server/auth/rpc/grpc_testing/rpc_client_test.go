/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package grpc_testing

import (
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"context"
	"testing"
	"fmt"
	"google.golang.org/grpc/metadata"
)

func TestRPCClient(t *testing.T)  {
	fmt.Println("TestRPCClient...")
	conn, err := grpc.Dial("127.0.0.1:10001", grpc.WithInsecure())
	if err != nil {
		glog.Fatalf("fail to dial: %v\n", err)
	}
	defer conn.Close()
	client := NewAuthClient(conn)
	authSendCode := &TLAuthSendCode{
		Flags: 1,
		AllowFlashcall: false,
		PhoneNumber: "+86 111 1111 1111",
		// CurrentNumber: &Bool{Payload:&TLBoolTrue{}},
		ApiId: 1,
		ApiHash: "1",
	}

	md := metadata.Pairs("key1", "val1_1", "key2", "val2", "key3", "val3", "key1", "val1_2")
	// create a new context with this metadata
	ctx := metadata.NewOutgoingContext(context.Background(), md)


	// glog.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	auth_SentCode, err := client.AuthSentCode(ctx, authSendCode)
	if err != nil {
		fmt.Errorf("%v.AuthSentCode(_) = _, %v: \n", client, err)
	}
	fmt.Printf("%v\n", auth_SentCode)
}
