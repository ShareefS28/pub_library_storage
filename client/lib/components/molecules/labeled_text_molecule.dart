import 'package:flutter/material.dart';
import 'package:library_storage/components/atoms/app_text_atom.dart';
import 'package:library_storage/components/atoms/atoms.dart';

class LabeledTextMolecule extends StatelessWidget {
  final String label;
  final String value;

  const LabeledTextMolecule({
    super.key,
    required this.label,
    required this.value
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        AppTextAtom(text: label, style: AppTextStylesAtom.caption),
        AppTextAtom(text: value, style: AppTextStylesAtom.body)
      ],
    );
  }

}