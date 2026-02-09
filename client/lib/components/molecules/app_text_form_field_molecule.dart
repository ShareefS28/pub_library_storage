import 'package:flutter/material.dart';
import 'package:library_storage/components/atoms/app_text_styles_atom.dart';
import '../atoms/atoms.dart';

class AppTextFormFieldMolecule extends StatefulWidget {
  final String label;
  final String? hint;
  final TextEditingController controller;
  final String? Function(String?)? validator;
  final TextFieldModeAtom mode;

  const AppTextFormFieldMolecule ({
    super.key,
    required this.label,
    required this.controller,
    required this.mode,
    this.hint,
    this.validator
  });

  @override
  State<AppTextFormFieldMolecule> createState() => _AppTextFormFieldMolecule();
}

class _AppTextFormFieldMolecule extends State<AppTextFormFieldMolecule> {

  bool _obsecureText = true;

  bool get isPassword => widget.mode == TextFieldModeAtom.password;

  TextInputType get keyboardType => switch (widget.mode) {
    TextFieldModeAtom.text => TextInputType.text,
    TextFieldModeAtom.email => TextInputType.emailAddress,
    TextFieldModeAtom.password => TextInputType.visiblePassword,
    TextFieldModeAtom.number => TextInputType.number,
    _ => TextInputType.text
  };

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(widget.label, style: AppTextStylesAtom.label),
        const SizedBox(height: 6),
        TextFormField(
          controller: widget.controller,
          validator: widget.validator,
          obscureText: isPassword ? _obsecureText : false,
          keyboardType: keyboardType,
          style: AppTextStylesAtom.input,
          decoration: InputDecoration(
            hintText: widget.hint,
            suffixIcon: isPassword ? IconButton(
              icon: Icon(
                _obsecureText ? Icons.visibility_off : Icons.visibility,
              ),
              onPressed: () {
                setState(() {
                  _obsecureText = !_obsecureText;
                });
              },
            ) : null,
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(8),
              borderSide: const BorderSide(color: AppColorsAtom.border),
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(8),
              borderSide: const BorderSide(color: AppColorsAtom.primary),
            ),
            errorBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(8),
              borderSide: const BorderSide(color: AppColorsAtom.error),
            ),
          ),
        ),
      ],
    );
  }
}