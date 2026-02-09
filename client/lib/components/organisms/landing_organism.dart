import 'package:flutter/material.dart';
import '../molecules/molecules.dart';

class LandingOrganism extends StatelessWidget {
  final bool isLoading;

  const LandingOrganism({
    super.key,
    required this.isLoading,
  });

  @override
  Widget build(BuildContext context) {
    return Center(
      child: isLoading ? const Center(child: CircularProgressIndicator()) : LabeledTextMolecule(label: "WELCOME!", value: "Landing Page")
    );
  }
}