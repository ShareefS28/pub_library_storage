import 'package:flutter/material.dart';
import '../organisms/organisms.dart';

class LandingTemplate extends StatelessWidget {
  final bool isLoading;
  final Widget sidebar;
  final Widget content;

  const LandingTemplate({ 
    super.key,
    required this.isLoading,
    required this.sidebar,
    required this.content
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [ 
        Container(
          width: MediaQuery.of(context).size.width * 0.9,
          height: MediaQuery.of(context).size.height * 0.9,
          padding: const EdgeInsets.all(8),
          decoration: BoxDecoration(
            border: Border.all(
              color: Colors.grey,
              width: 1.5
            ),
            borderRadius: BorderRadius.circular(12)
          ),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Expanded(
                child: Column(
                  children: [
                    sidebar,
                    Expanded(
                      child: isLoading ? const Center(child: CircularProgressIndicator()) : content,
                    )
                  ],
                )
              )
            ]
          )
        )
      ]
    );  
  }

}