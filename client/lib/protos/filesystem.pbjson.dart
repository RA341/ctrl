//
//  Generated code. Do not modify.
//  source: protos/filesystem.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use inputFoldersDescriptor instead')
const InputFolders$json = {
  '1': 'InputFolders',
  '2': [
    {'1': 'srcPath', '3': 1, '4': 2, '5': 9, '10': 'srcPath'},
    {'1': 'destPath', '3': 2, '4': 2, '5': 9, '10': 'destPath'},
  ],
};

/// Descriptor for `InputFolders`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List inputFoldersDescriptor = $convert.base64Decode(
    'CgxJbnB1dEZvbGRlcnMSGAoHc3JjUGF0aBgBIAIoCVIHc3JjUGF0aBIaCghkZXN0UGF0aBgCIA'
    'IoCVIIZGVzdFBhdGg=');

@$core.Deprecated('Use linkResultDescriptor instead')
const LinkResult$json = {
  '1': 'LinkResult',
  '2': [
    {'1': 'error', '3': 1, '4': 1, '5': 9, '10': 'error'},
  ],
};

/// Descriptor for `LinkResult`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List linkResultDescriptor = $convert.base64Decode(
    'CgpMaW5rUmVzdWx0EhQKBWVycm9yGAEgASgJUgVlcnJvcg==');

@$core.Deprecated('Use pathDescriptor instead')
const Path$json = {
  '1': 'Path',
  '2': [
    {'1': 'path', '3': 1, '4': 2, '5': 9, '10': 'path'},
  ],
};

/// Descriptor for `Path`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pathDescriptor = $convert.base64Decode(
    'CgRQYXRoEhIKBHBhdGgYASACKAlSBHBhdGg=');

@$core.Deprecated('Use folderDescriptor instead')
const Folder$json = {
  '1': 'Folder',
  '2': [
    {'1': 'fullPath', '3': 1, '4': 2, '5': 9, '10': 'fullPath'},
    {'1': 'files', '3': 2, '4': 3, '5': 11, '6': '.fs.File', '10': 'files'},
    {'1': 'folders', '3': 3, '4': 3, '5': 11, '6': '.fs.Folder', '10': 'folders'},
  ],
};

/// Descriptor for `Folder`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List folderDescriptor = $convert.base64Decode(
    'CgZGb2xkZXISGgoIZnVsbFBhdGgYASACKAlSCGZ1bGxQYXRoEh4KBWZpbGVzGAIgAygLMgguZn'
    'MuRmlsZVIFZmlsZXMSJAoHZm9sZGVycxgDIAMoCzIKLmZzLkZvbGRlclIHZm9sZGVycw==');

@$core.Deprecated('Use fileDescriptor instead')
const File$json = {
  '1': 'File',
  '2': [
    {'1': 'name', '3': 1, '4': 2, '5': 9, '10': 'name'},
  ],
};

/// Descriptor for `File`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List fileDescriptor = $convert.base64Decode(
    'CgRGaWxlEhIKBG5hbWUYASACKAlSBG5hbWU=');

