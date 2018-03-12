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

package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/mtproto"
	"net"
	"google.golang.org/grpc"
	account "github.com/nebulaim/telegramd/biz_server/account/rpc"
	auth "github.com/nebulaim/telegramd/biz_server/auth/rpc"
	bots "github.com/nebulaim/telegramd/biz_server/bots/rpc"
	channels "github.com/nebulaim/telegramd/biz_server/channels/rpc"
	contacts "github.com/nebulaim/telegramd/biz_server/contacts/rpc"
	help "github.com/nebulaim/telegramd/biz_server/help/rpc"
	langpack "github.com/nebulaim/telegramd/biz_server/langpack/rpc"
	messages "github.com/nebulaim/telegramd/biz_server/messages/rpc"
	payments "github.com/nebulaim/telegramd/biz_server/payments/rpc"
	phone "github.com/nebulaim/telegramd/biz_server/phone/rpc"
	photos "github.com/nebulaim/telegramd/biz_server/photos/rpc"
	stickers "github.com/nebulaim/telegramd/biz_server/stickers/rpc"
	updates "github.com/nebulaim/telegramd/biz_server/updates/rpc"
	upload "github.com/nebulaim/telegramd/biz_server/upload/rpc"
	users "github.com/nebulaim/telegramd/biz_server/users/rpc"
	"github.com/nebulaim/telegramd/base/orm"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")


	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(err)
	}

	// TODO(@benqi): 从配置文件载入
	err = orm.RegisterDataBase("default", "mysql",  "root:@/nebulaim?charset=utf8", 30)
	if err != nil {
		panic(err)
	}
}

// 整合各服务，方便开发调试
func main() {
	flag.Parse()

	zorm := orm.NewOrm()
	if zorm != nil {
		glog.Fatal("failed to connect mysql database, dsn: root:@/nebulaim?charset=utf8")
	}

	lis, err := net.Listen("tcp", "0.0.0.0:10001")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	mtproto.RegisterRPCAccountServer(grpcServer, &account.AccountServiceImpl{zorm})
	mtproto.RegisterRPCAuthServer(grpcServer, &auth.AuthServiceImpl{zorm})
	mtproto.RegisterRPCBotsServer(grpcServer, &bots.BotsServiceImpl{zorm})
	mtproto.RegisterRPCChannelsServer(grpcServer, &channels.ChannelsServiceImpl{zorm})
	mtproto.RegisterRPCContactsServer(grpcServer, &contacts.ContactsServiceImpl{zorm})
	mtproto.RegisterRPCHelpServer(grpcServer, &help.HelpServiceImpl{zorm})
	mtproto.RegisterRPCLangpackServer(grpcServer, &langpack.LangpackServiceImpl{zorm})
	mtproto.RegisterRPCMessagesServer(grpcServer, &messages.MessagesServiceImpl{zorm})
	mtproto.RegisterRPCPaymentsServer(grpcServer, &payments.PaymentsServiceImpl{zorm})
	mtproto.RegisterRPCPhoneServer(grpcServer, &phone.PhoneServiceImpl{zorm})
	mtproto.RegisterRPCPhotosServer(grpcServer, &photos.PhotosServiceImpl{zorm})
	mtproto.RegisterRPCStickersServer(grpcServer, &stickers.StickersServiceImpl{zorm})
	mtproto.RegisterRPCUpdatesServer(grpcServer, &updates.UpdatesServiceImpl{zorm})
	mtproto.RegisterRPCUploadServer(grpcServer, &upload.UploadServiceImpl{zorm})
	mtproto.RegisterRPCUsersServer(grpcServer, &users.UsersServiceImpl{zorm})

	glog.Info("NewRPCServer in 0.0.0.0:10001.")

	grpcServer.Serve(lis)
}
