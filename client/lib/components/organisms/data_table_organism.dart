import 'package:flutter/material.dart';
import '../molecules/molecules.dart';

class DataTableOrganism extends StatelessWidget {
  final List<DataColumn> columns;
  final List<DataRowMolecule> rows;

  const DataTableOrganism({
    super.key,
    required this.columns,
    required this.rows
  });

  @override
  Widget build(BuildContext context) {
    return DataTable(
      columns: this.columns,
      rows: this.rows.map((row) => row.build()).toList(),
    );
  }

}