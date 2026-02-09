import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import 'package:library_storage/services/ApiClient.dart';
import 'package:library_storage/utils/globals.dart';
import 'components/pages/pages.dart';

void main() {
  Apiclient.setupInterceptors();

  final globalData = GlobalData();
  globalData.checkAuth(); // start auth check

  runApp(
    ChangeNotifierProvider.value(
      value: globalData,
      child: const MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    final globalData = Provider.of<GlobalData>(context, listen: true);

    final router = GoRouter(
      initialLocation: '/',
      refreshListenable: globalData,
      routes: [
        GoRoute(
          path: '/',
          builder: (context, state) => const LandingPage(),
        ),
        GoRoute(
          path: '/login',
          builder: (context, state) => const LoginPage(),
        ),
        GoRoute(
          path: '/register',
          builder: (context, state) => const RegisterPage(),
        ),
      ],
      redirect: (context, state) {
        final isLoggedIn = globalData.isLoggedIn;
        final isChecking = globalData.isCheckingAuth;

        // Show splash/loading while checking auth
        if (isChecking) return null; // stay on current page or show a loader

        // Current path
        final currentPath = state.uri.path;

        // Routes that don't require auth
        final loggingIn = currentPath == '/login' || currentPath == '/register';

        // Not logged in → send to login
        if (!isLoggedIn && !loggingIn) return '/login';

        // Logged in → prevent going to login/register
        if (isLoggedIn && loggingIn) return '/';

        return null; // no redirect
      },
    );

    return MaterialApp.router(
      debugShowCheckedModeBanner: true,
      routerConfig: router,
    );
  }
}
