import 'package:flutter/material.dart';
import 'app_colors_atom.dart';

class AppButtonStylesAtom {
  static ButtonStyle primary = ElevatedButton.styleFrom(
    backgroundColor: AppColorsAtom.primary,
    foregroundColor: Colors.white,
    padding: const EdgeInsets.symmetric(vertical: 14),
    shape: RoundedRectangleBorder(
      borderRadius: BorderRadiusGeometry.directional(
        topStart: Radius.circular(8),
        topEnd: Radius.circular(8),
        bottomStart: Radius.circular(8),  
        bottomEnd:  Radius.circular(8)
      )
    ),
  );

  static ButtonStyle cancel = ElevatedButton.styleFrom(
    backgroundColor: AppColorsAtom.error,
    foregroundColor: Colors.white,
    padding: const EdgeInsets.symmetric(vertical: 14),
    shape: RoundedRectangleBorder(
      borderRadius: BorderRadiusGeometry.directional(
        topStart: Radius.circular(8),
        topEnd: Radius.circular(8),
        bottomStart: Radius.circular(8),  
        bottomEnd:  Radius.circular(8)
      )
    ),
  );

  static ButtonStyle confirm = ElevatedButton.styleFrom(
    backgroundColor: AppColorsAtom.confirm,
    foregroundColor: Colors.white,
    padding: const EdgeInsets.symmetric(vertical: 14),
    shape: RoundedRectangleBorder(
      borderRadius: BorderRadiusGeometry.directional(
        topStart: Radius.circular(8),
        topEnd: Radius.circular(8),
        bottomStart: Radius.circular(8),  
        bottomEnd:  Radius.circular(8)
      )
    ),
  );
}