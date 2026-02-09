import 'package:flutter/material.dart';
import '../../../components/atoms/atoms.dart';
import '../../../components/molecules/molecules.dart';

class BookTableDto {
  final String title;

  BookTableDto({
    required this.title
  });

  // Convert JSON to Dart object
  factory BookTableDto.fromJson(Map<String, dynamic> json) {
    final items = json['data'];

    return BookTableDto(
      title: items['title'],
    );
  }
}