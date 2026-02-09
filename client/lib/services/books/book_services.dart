import 'package:dio/dio.dart';
import 'package:library_storage/services/ApiClient.dart';
import '../../utils/utils.dart';

class BookServices {
  static final Dio dio = Apiclient.dio;

  // Post Login
  static Future<dynamic> fetchBooks() async {
    try {
      final response = await dio.get('/books');

      return response;
    } on DioException catch (e) {
      throw Exception('Login Failed: ${e.response?.statusCode} ${e.message}');
    }
  }
}