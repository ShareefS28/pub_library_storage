import 'package:flutter/material.dart';
import '../organisms/organisms.dart';

class LoginTemplate extends StatelessWidget {
  final GlobalKey<FormState> formKey;
  final TextEditingController usernameController;
  final TextEditingController passwordController;
  final bool isLoading;
  final VoidCallback onSubmit;
  final VoidCallback onRegister;

  const LoginTemplate({ 
    super.key,
    required this.formKey,
    required this.usernameController,
    required this.passwordController,
    required this.isLoading,
    required this.onSubmit,
    required this.onRegister,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [ 
        Container(
          width: MediaQuery.of(context).size.width * 0.5,
          padding: const EdgeInsets.all(16),
          decoration: BoxDecoration(
            border: Border.all(
              color: Colors.grey,
              width: 1.5
            ),
            borderRadius: BorderRadius.circular(12)
          ),
          child: LoginOrganism(
            formKey: formKey,
            usernameController: usernameController,
            passwordController: passwordController,
            isLoading: isLoading,
            onSubmit: onSubmit,
            onRegister: onRegister,
          ),
        )
      ]
    );
  }
}