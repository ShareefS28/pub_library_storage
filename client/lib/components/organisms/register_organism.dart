import 'package:flutter/material.dart';
import '../atoms/atoms.dart';
import '../molecules/molecules.dart';
import '../../utils/validators/validators.dart';

class RegisterOrganism extends StatelessWidget {
  final GlobalKey<FormState> formKey;
  final TextEditingController usernameController;
  final TextEditingController passwordController;
  final TextEditingController confirmPasswordController;
  final bool isLoading;
  final VoidCallback onSubmit;

  const RegisterOrganism({ 
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
          const SizedBox(height: 6),
          AppTextFormFieldMolecule(
            mode: TextFieldModeAtom.password,
            label: "Confirm Password",
            hint: "Enter your password",
            controller: confirmPasswordController,
            validator: (value) => Validators.confirmPassword(value, passwordController.text)
          ),
          const SizedBox(height: 24),
          AppButtonMolecule(
            text: "Register",
            style: AppButtonStylesAtom.confirm,
            isLoading: isLoading, 
            onPressed: onSubmit
          )
        ],
      ),
    );
  }
}