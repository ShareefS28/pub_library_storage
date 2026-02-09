import 'package:cookie_jar/cookie_jar.dart';
import 'package:dio/browser.dart';
import 'package:dio/dio.dart';
import 'package:dio_cookie_manager/dio_cookie_manager.dart';
import 'package:flutter/foundation.dart';
import 'package:library_storage/utils/utils.dart';

class Apiclient {
  static final Dio dio = Dio(
    BaseOptions(
      baseUrl: Apiconstant.baseUrl,
      connectTimeout: Apiconstant.connectTimeout,
      receiveTimeout: Apiconstant.receiveTimeout,
      headers: {
        'Content-Type': 'application/json'
      },
      validateStatus: (status) => true,
    ),
  );

  static void setupInterceptors() {
    if (!kIsWeb) {
      // Only add CookieManager on mobile
      dio.interceptors.add(CookieManager(CookieJar()));
    } else{
      dio.httpClientAdapter = BrowserHttpClientAdapter()..withCredentials = true;
    }
  }
}