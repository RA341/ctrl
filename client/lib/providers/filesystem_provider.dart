import 'dart:async';

import 'package:client/protos/filesystem.pb.dart';
import 'package:client/services/hello.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final getFolderProvider =
    FutureProvider.family<Folder, Path>((ref, Path path) async {
  return FsService.i.client.listFiles(path);
});

final folderProvider = AutoDisposeAsyncNotifierProviderFamily<FileFetcher, Folder, int>(
  () => FileFetcher(),
);

class FileFetcher extends AutoDisposeFamilyAsyncNotifier<Folder, int> {
  FileFetcher();

  @override
  FutureOr<Folder> build(int arg) {
    return ref.watch(getFolderProvider(Path(path: '.')).future);
  }

  Future<void> changeDir(String path) async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(
      () async {
        return await ref.watch(
          getFolderProvider(Path(path: path)).future,
        );
      },
    );
  }
}
