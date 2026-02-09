import 'package:flutter/material.dart';
import '../atoms/atoms.dart';
import '../molecules/molecules.dart';
import '../organisms/organisms.dart';
import '../templates/templates.dart';
import '../../services/services.dart';
import '../../utils/DTOs/dto.dart';
import '../templates/login_template.dart';

class LandingPage extends StatefulWidget {
  const LandingPage({ super.key });

  @override
  State<LandingPage> createState() => _LandingPageState();
}

class _LandingPageState extends State<LandingPage> {
  bool isLoading = false;

  final List<BookTableDto> mockBooks = List.generate(
    30, 
    (index) => BookTableDto(title: 'หนังสือ ${index + 1}'),
  );

  List<DataColumn> buildColumns() {
    return const [
      DataColumn(
        label: AppTextAtom(text: 'ชื่อหนังสือ')
      ),
    ];
  }

  List<DataRowMolecule> buildRows(List<BookTableDto> books) {
    return books.map((book) {
      return DataRowMolecule([
        AppTextAtom(text: book.title)
      ]);
    }).toList();
  }

  @override
  void initState(){
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Landing Page"),
        centerTitle: true,
      ),
      body: SafeArea(
        child: LandingTemplate(
          isLoading: isLoading,
          sidebar: LandingOrganism(isLoading: isLoading),
          content: DataTableOrganism(
            columns: buildColumns(), 
            rows: buildRows(mockBooks)
          ),
        )
      ),
    );
  }

  @override
  void dispose() {
    super.dispose();
  }
}