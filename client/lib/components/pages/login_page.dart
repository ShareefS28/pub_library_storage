import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:library_storage/services/login/login_services.dart';
import 'package:library_storage/utils/DTOs/dto.dart';
import 'package:library_storage/utils/globals.dart';
import 'package:provider/provider.dart';
import '../templates/login_template.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({ super.key });

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _formKey = GlobalKey<FormState>();
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();
  bool isLoading = false;

  void handleLogin() async {
    if (!_formKey.currentState!.validate()) return;

    // Validate Fields
    if (usernameController.text.isEmpty || passwordController.text.isEmpty) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Please fill all fields.')),
      );

      return;
    } 

    setState(() {
      isLoading = true;
    });

    // Create DTO here (UI -> Data layer)
    final dto = LoginRequestDto(
      username: usernameController.text.trim(),
      password: passwordController.text
    );

    // API call
    try {
      final response = await LoginServices.login(dto);

      if (response.statusCode == 200) {
        var globalData = Provider.of<GlobalData>(context, listen: false);
        // final loginResponse = LoginResponseDto.fromJson(response.data);
        globalData.login();
      }
      else {
        final errResponse = ErrorResponseDto.fromJson(response.data);
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Register failed ${errResponse.message}')),
        );
      }

    } catch (e) {
      print(e);
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Login failed ${e}')),
      );
    } finally {
      setState(() {
        isLoading = false;
      });
    }
  }

  @override
  void initState(){
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Login"),
        centerTitle: true,
      ),
      body: SafeArea(
        child: LoginTemplate(
          formKey: _formKey,
          usernameController: usernameController, 
          passwordController: passwordController, 
          isLoading: isLoading, 
          onSubmit: handleLogin,
          onRegister: ()  {
            context.push('/register');
          },
        )
      ),
    );
  }

  @override
  void dispose() {
    usernameController.dispose();
    passwordController.dispose();
    super.dispose();
  }
}