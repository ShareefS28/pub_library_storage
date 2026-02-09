import 'package:flutter/material.dart';
import 'package:library_storage/components/atoms/app_text_atom.dart';
import 'package:library_storage/components/atoms/app_text_styles_atom.dart';
import '../atoms/app_button_styles_atom.dart';

class AppButtonMolecule extends StatelessWidget {
  final String text;
  final ButtonStyle? style;
  final VoidCallback? onPressed;
  final bool isLoading;

  const AppButtonMolecule({
    super.key,
    required this.text,
    this.style,
    required this.onPressed,
    this.isLoading = false,
  });

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: ElevatedButton(
        style: style ?? AppButtonStylesAtom.primary,
        onPressed: isLoading ? null : onPressed,
        child: isLoading ? const SizedBox(
          height: 22,
          width: 22,
          child: CircularProgressIndicator(
            strokeWidth: 2,
            color: Colors.white,
          ),
        ) : AppTextAtom(
          text: text,
          style: AppTextStylesAtom.label,
          textAlign: TextAlign.center,
        )
      ),
    );
  }
}

