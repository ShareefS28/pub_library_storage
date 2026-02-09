import 'package:flutter/material.dart';

class DataCellAtom {
  final Widget child;

  DataCellAtom(this.child);

  DataCell build() {
    return DataCell(child);
  }

}
