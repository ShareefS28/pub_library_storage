import 'package:flutter/material.dart';

class AppTextAtom extends StatelessWidget {
  final String text;
  final TextStyle? style;
  final TextOverflow overflow;
  final int? maxLines;
  final TextAlign? textAlign;
  final bool isLoading;

  const AppTextAtom({
    super.key,
    required this.text,
    this.style,
    this.isLoading = false,
    this.overflow = TextOverflow.fade,
    this.maxLines,
    this.textAlign,
  });

  @override
  Widget build(BuildContext context) {
    if (isLoading) {
      return const SizedBox(
        height: 14,
        width: 80,
        child: LinearProgressIndicator(),
      );
    }

    return Text(
      text,
      style: style,
      overflow: overflow,
      maxLines: maxLines,
      textAlign: textAlign,
    );
  }
}

