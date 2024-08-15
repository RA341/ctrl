//
//  Generated code. Do not modify.
//  source: protos/filesystem.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class InputFolders extends $pb.GeneratedMessage {
  factory InputFolders({
    $core.String? srcPath,
    $core.String? destPath,
  }) {
    final $result = create();
    if (srcPath != null) {
      $result.srcPath = srcPath;
    }
    if (destPath != null) {
      $result.destPath = destPath;
    }
    return $result;
  }
  InputFolders._() : super();
  factory InputFolders.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory InputFolders.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'InputFolders', package: const $pb.PackageName(_omitMessageNames ? '' : 'fs'), createEmptyInstance: create)
    ..aQS(1, _omitFieldNames ? '' : 'srcPath', protoName: 'srcPath')
    ..aQS(2, _omitFieldNames ? '' : 'destPath', protoName: 'destPath')
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  InputFolders clone() => InputFolders()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  InputFolders copyWith(void Function(InputFolders) updates) => super.copyWith((message) => updates(message as InputFolders)) as InputFolders;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static InputFolders create() => InputFolders._();
  InputFolders createEmptyInstance() => create();
  static $pb.PbList<InputFolders> createRepeated() => $pb.PbList<InputFolders>();
  @$core.pragma('dart2js:noInline')
  static InputFolders getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<InputFolders>(create);
  static InputFolders? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get srcPath => $_getSZ(0);
  @$pb.TagNumber(1)
  set srcPath($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSrcPath() => $_has(0);
  @$pb.TagNumber(1)
  void clearSrcPath() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get destPath => $_getSZ(1);
  @$pb.TagNumber(2)
  set destPath($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDestPath() => $_has(1);
  @$pb.TagNumber(2)
  void clearDestPath() => clearField(2);
}

class LinkResult extends $pb.GeneratedMessage {
  factory LinkResult({
    $core.String? error,
  }) {
    final $result = create();
    if (error != null) {
      $result.error = error;
    }
    return $result;
  }
  LinkResult._() : super();
  factory LinkResult.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LinkResult.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'LinkResult', package: const $pb.PackageName(_omitMessageNames ? '' : 'fs'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'error')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  LinkResult clone() => LinkResult()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  LinkResult copyWith(void Function(LinkResult) updates) => super.copyWith((message) => updates(message as LinkResult)) as LinkResult;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static LinkResult create() => LinkResult._();
  LinkResult createEmptyInstance() => create();
  static $pb.PbList<LinkResult> createRepeated() => $pb.PbList<LinkResult>();
  @$core.pragma('dart2js:noInline')
  static LinkResult getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LinkResult>(create);
  static LinkResult? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get error => $_getSZ(0);
  @$pb.TagNumber(1)
  set error($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasError() => $_has(0);
  @$pb.TagNumber(1)
  void clearError() => clearField(1);
}

class Path extends $pb.GeneratedMessage {
  factory Path({
    $core.String? path,
  }) {
    final $result = create();
    if (path != null) {
      $result.path = path;
    }
    return $result;
  }
  Path._() : super();
  factory Path.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Path.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Path', package: const $pb.PackageName(_omitMessageNames ? '' : 'fs'), createEmptyInstance: create)
    ..aQS(1, _omitFieldNames ? '' : 'path')
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Path clone() => Path()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Path copyWith(void Function(Path) updates) => super.copyWith((message) => updates(message as Path)) as Path;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Path create() => Path._();
  Path createEmptyInstance() => create();
  static $pb.PbList<Path> createRepeated() => $pb.PbList<Path>();
  @$core.pragma('dart2js:noInline')
  static Path getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Path>(create);
  static Path? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get path => $_getSZ(0);
  @$pb.TagNumber(1)
  set path($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasPath() => $_has(0);
  @$pb.TagNumber(1)
  void clearPath() => clearField(1);
}

class Folder extends $pb.GeneratedMessage {
  factory Folder({
    $core.String? fullPath,
    $core.Iterable<File>? files,
    $core.Iterable<Folder>? folders,
  }) {
    final $result = create();
    if (fullPath != null) {
      $result.fullPath = fullPath;
    }
    if (files != null) {
      $result.files.addAll(files);
    }
    if (folders != null) {
      $result.folders.addAll(folders);
    }
    return $result;
  }
  Folder._() : super();
  factory Folder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Folder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Folder', package: const $pb.PackageName(_omitMessageNames ? '' : 'fs'), createEmptyInstance: create)
    ..aQS(1, _omitFieldNames ? '' : 'fullPath', protoName: 'fullPath')
    ..pc<File>(2, _omitFieldNames ? '' : 'files', $pb.PbFieldType.PM, subBuilder: File.create)
    ..pc<Folder>(3, _omitFieldNames ? '' : 'folders', $pb.PbFieldType.PM, subBuilder: Folder.create)
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Folder clone() => Folder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Folder copyWith(void Function(Folder) updates) => super.copyWith((message) => updates(message as Folder)) as Folder;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Folder create() => Folder._();
  Folder createEmptyInstance() => create();
  static $pb.PbList<Folder> createRepeated() => $pb.PbList<Folder>();
  @$core.pragma('dart2js:noInline')
  static Folder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Folder>(create);
  static Folder? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get fullPath => $_getSZ(0);
  @$pb.TagNumber(1)
  set fullPath($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFullPath() => $_has(0);
  @$pb.TagNumber(1)
  void clearFullPath() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<File> get files => $_getList(1);

  @$pb.TagNumber(3)
  $core.List<Folder> get folders => $_getList(2);
}

class File extends $pb.GeneratedMessage {
  factory File({
    $core.String? name,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    return $result;
  }
  File._() : super();
  factory File.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory File.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'File', package: const $pb.PackageName(_omitMessageNames ? '' : 'fs'), createEmptyInstance: create)
    ..aQS(1, _omitFieldNames ? '' : 'name')
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  File clone() => File()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  File copyWith(void Function(File) updates) => super.copyWith((message) => updates(message as File)) as File;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static File create() => File._();
  File createEmptyInstance() => create();
  static $pb.PbList<File> createRepeated() => $pb.PbList<File>();
  @$core.pragma('dart2js:noInline')
  static File getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<File>(create);
  static File? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
