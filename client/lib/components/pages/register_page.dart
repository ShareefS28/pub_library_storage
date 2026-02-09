import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:library_storage/services/login/login_services.dart';
import 'package:library_storage/services/register/register_services.dart';
import 'package:library_storage/utils/DTOs/Error/error_response.dart';
import 'package:library_storage/utils/DTOs/dto.dart';
import '../templates/templates.dart';

class RegisterPage extends StatefulWidget {
  const RegisterPage({ super.key });

  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final _formKey = GlobalKey<FormState>();
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();
  final confirmPasswordController = TextEditingController();
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
    final dto = RegisterRequestDto(
      username: usernameController.text.trim(),
      password: passwordController.text
    );

    // API call
    try {
      final response = await RegisterServices.Register(dto);

      if (response.statusCode == 200) {
        context.pop('/login');
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
        child: RegisterTemplate(
          formKey: _formKey,
          usernameController: usernameController, 
          passwordController: passwordController, 
          confirmPasswordController: confirmPasswordController,
          isLoading: isLoading, 
          onSubmit: handleLogin
        )
      ),
    );
  }

  @override
  void dispose() {
    usernameController.dispose();
    passwordController.dispose();
    confirmPasswordController.dispose();
    super.dispose();
  }
}