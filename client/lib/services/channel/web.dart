import 'package:grpc/grpc_web.dart';

GrpcWebClientChannel createChannel(String baseUrl, String port) {
  print('init web');

  return GrpcWebClientChannel.xhr(Uri.parse('http://$baseUrl:$port'));
}
