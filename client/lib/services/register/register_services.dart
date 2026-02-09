import 'package:dio/dio.dart';
import 'package:library_storage/services/ApiClient.dart';
import '../../utils/utils.dart';

class RegisterServices {
  static final Dio dio = Apiclient.dio;

  // Post Register
  static Future<dynamic> Register(RegisterRequestDto data) async {
    try {
      final response = await dio.post(
        '/auth/Register',
        data: data.toJson(),
      );

      return response;
    } on DioException catch (e) {
      throw Exception('Register Failed: ${e.response?.statusCode} ${e.message}');
    }
  }
}