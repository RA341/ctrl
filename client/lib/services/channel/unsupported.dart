import 'package:grpc/grpc.dart';

Never _unsupported() {
  throw UnsupportedError(
      'No suitable database implementation was found on this platform.');
}

ClientChannel createChannel(String baseUrl, String port) {
  throw UnsupportedError(
      'No suitable database implementation was found on this platform.');
}
