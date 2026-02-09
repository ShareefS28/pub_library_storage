import 'package:flutter/material.dart';
import '../atoms/atoms.dart';
import '../molecules/molecules.dart';
import '../../utils/validators/validators.dart';

class LoginOrganism extends StatelessWidget {
  final GlobalKey<FormState> formKey;
  final TextEditingController usernameController;
  final TextEditingController passwordController;
  final bool isLoading;
  final VoidCallback onSubmit;
  final VoidCallback onRegister;

  const LoginOrganism({ 
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
    return Form(
      key: formKey,
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          AppTextFormFieldMolecule(
            mode: TextFieldModeAtom.text,
            label: "Username",
            hint: "Enter your username",
            controller: usernameController, 
            validator: (value) => Validators.required(value, fieldName: "Username")
          ),
          const SizedBox(height: 6),
          AppTextFormFieldMolecule(
            mode: TextFieldModeAtom.password,
            label: "Password",
            hint: "Enter your password",
            controller: passwordController,
            validator: (value) => Validators.password(value)
          ),
          const SizedBox(height: 24),
          AppButtonMolecule(
            text: "Login",
            style: AppButtonStylesAtom.confirm,
            isLoading: isLoading, 
            onPressed: onSubmit
          ),
          const SizedBox(height: 6),
          AppButtonMolecule(
            text: "Register",
            style: AppButtonStylesAtom.primary,
            onPressed: onRegister
          )
        ],
      ),
    );
  }
}