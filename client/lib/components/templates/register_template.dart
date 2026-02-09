import 'package:flutter/material.dart';
import 'package:library_storage/components/organisms/register_organism.dart';
import '../organisms/organisms.dart';

class RegisterTemplate extends StatelessWidget {
  final GlobalKey<FormState> formKey;
  final TextEditingController usernameController;
  final TextEditingController passwordController;
  final TextEditingController confirmPasswordController;
  final bool isLoading;
  final VoidCallback onSubmit;

  const RegisterTemplate({ 
    super.key,
    required this.formKey,
    required this.usernameController,
    required this.passwordController,
    required this.confirmPasswordController,
    required this.isLoading,
    required this.onSubmit,
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
          child: RegisterOrganism(
            formKey: formKey,
            usernameController: usernameController,
            passwordController: passwordController,
            confirmPasswordController: confirmPasswordController,
            isLoading: isLoading,
            onSubmit: onSubmit
          ),
        )
      ]
    );
  }
}