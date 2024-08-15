import 'package:grpc/grpc.dart';

ClientChannel createChannel(String baseUrl, String port) {
  print('init native');
  return ClientChannel(
    baseUrl,
    port: int.parse(port),
    options: const ChannelOptions(credentials: ChannelCredentials.insecure()),
  );
}
