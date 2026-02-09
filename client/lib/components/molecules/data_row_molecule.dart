import 'package:flutter/material.dart';
import '../atoms/atoms.dart';

class DataRowMolecule {
  final List<Widget> cells;

  DataRowMolecule(this.cells);

  DataRow build() {
    return DataRow(
      cells: cells.map((cell) => DataCellAtom(cell).build()).toList()
    );
  }
}