import 'package:client/protos/filesystem.pb.dart';
import 'package:client/providers/filesystem_provider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:path/path.dart' as path;

class FolderView extends ConsumerWidget {
  const FolderView({
    required this.tabIndex,
    super.key,
  });

  final int tabIndex;

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final folder = ref.watch(folderProvider(tabIndex));

    return folder.when(
      data: (data) {
        final folderList = data.folders;
        final fileList = data.files;

        return Column(
          children: [
            AppBar(
              leading: IconButton(
                onPressed: () {
                  ref
                      .read(folderProvider(tabIndex).notifier)
                      .changeDir(path.dirname(data.fullPath));
                },
                icon: const Icon(Icons.arrow_back),
              ),
              backgroundColor: Theme.of(context).colorScheme.inversePrimary,
              title: Text(data.fullPath),
            ),
            Expanded(
              child: ListView.builder(
                itemCount: folderList.length + fileList.length,
                itemBuilder: (context, index) {
                  if (folderList.length > index) {
                    final folder = folderList[index];
                    return FolderTile(folder: folder, tabIndex: tabIndex);
                  }

                  final file = fileList.elementAt(
                    index - folderList.length,
                  );
                  return FileTile(file: file);
                },
              ),
            ),
          ],
        );
      },
      error: (error, stackTrace) {
        return Text(error.toString());
      },
      loading: () => const CircularProgressIndicator(),
    );
  }
}

class FolderTile extends ConsumerWidget {
  const FolderTile({
    super.key,
    required this.folder,
    required this.tabIndex,
  });

  final Folder folder;
  final int tabIndex;

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return ListTile(
      leading: const Icon(Icons.folder),
      title: Text(path.basename(folder.fullPath)),
      onTap: () async {
        ref.read(folderProvider(tabIndex).notifier).changeDir(folder.fullPath);
      },
    );
  }
}

class FileTile extends ConsumerWidget {
  const FileTile({
    required this.file,
    super.key,
  });

  final File file;

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return ListTile(
      leading: const Icon(Icons.file_present),
      title: Text(file.name),
    );
  }
}
