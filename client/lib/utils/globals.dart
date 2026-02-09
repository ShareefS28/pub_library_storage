import 'package:flutter/material.dart';
import 'package:library_storage/services/ApiClient.dart';

class GlobalData extends ChangeNotifier {
  bool isLoggedIn = false;
  bool isCheckingAuth = true;

  void login() {
    isLoggedIn = true;
    notifyListeners();
  }

  void logout() {
    isLoggedIn = false;
    notifyListeners();
  }

  Future<void> checkAuth() async {
    try {
      final res = await Apiclient.dio.get('/secure/me');
      if (res.statusCode == 200) {
        isLoggedIn = true;
      }
    } catch (_) {
      isLoggedIn = false;
    }

    isCheckingAuth = false;
    notifyListeners();
  }
}
