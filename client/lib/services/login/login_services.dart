import 'package:dio/dio.dart';
import 'package:library_storage/services/ApiClient.dart';
import '../../utils/utils.dart';

class LoginServices {
  static final Dio dio = Apiclient.dio;

  // Post Login
  static Future<dynamic> login(LoginRequestDto data) async {
    try {
      final response = await dio.post(
        '/auth/login',
        data: data.toJson(),
      );

      return response;
    } on DioException catch (e) {
      throw Exception('Login Failed: ${e.response?.statusCode} ${e.message}');
    }
  }
}