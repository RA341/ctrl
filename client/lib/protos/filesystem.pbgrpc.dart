//
//  Generated code. Do not modify.
//  source: protos/filesystem.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'filesystem.pb.dart' as $0;

export 'filesystem.pb.dart';

@$pb.GrpcServiceName('fs.Filesystem')
class FilesystemClient extends $grpc.Client {
  static final _$listFiles = $grpc.ClientMethod<$0.Path, $0.Folder>(
      '/fs.Filesystem/ListFiles',
      ($0.Path value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Folder.fromBuffer(value));
  static final _$linkFolder = $grpc.ClientMethod<$0.InputFolders, $0.LinkResult>(
      '/fs.Filesystem/LinkFolder',
      ($0.InputFolders value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.LinkResult.fromBuffer(value));

  FilesystemClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.Folder> listFiles($0.Path request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$listFiles, request, options: options);
  }

  $grpc.ResponseFuture<$0.LinkResult> linkFolder($0.InputFolders request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$linkFolder, request, options: options);
  }
}

@$pb.GrpcServiceName('fs.Filesystem')
abstract class FilesystemServiceBase extends $grpc.Service {
  $core.String get $name => 'fs.Filesystem';

  FilesystemServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Path, $0.Folder>(
        'ListFiles',
        listFiles_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Path.fromBuffer(value),
        ($0.Folder value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.InputFolders, $0.LinkResult>(
        'LinkFolder',
        linkFolder_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.InputFolders.fromBuffer(value),
        ($0.LinkResult value) => value.writeToBuffer()));
  }

  $async.Future<$0.Folder> listFiles_Pre($grpc.ServiceCall call, $async.Future<$0.Path> request) async {
    return listFiles(call, await request);
  }

  $async.Future<$0.LinkResult> linkFolder_Pre($grpc.ServiceCall call, $async.Future<$0.InputFolders> request) async {
    return linkFolder(call, await request);
  }

  $async.Future<$0.Folder> listFiles($grpc.ServiceCall call, $0.Path request);
  $async.Future<$0.LinkResult> linkFolder($grpc.ServiceCall call, $0.InputFolders request);
}
