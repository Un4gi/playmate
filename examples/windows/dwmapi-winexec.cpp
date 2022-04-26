#include <iostream>
#include <windows.h>
#include <stdio.h>

#pragma comment(lib,"User32.lib")

#define DllExport __declspec(dllexport)

using namespace std;

DllExport void __stdcall SpawnCalc(void)
{
	WinExec("calc", 0);
}

BOOLEAN WINAPI DllMain(IN HINSTANCE hDllHandle, IN DWORD nReason, IN LPVOID Reserved)
{
   switch (nReason)
   {
   case DLL_PROCESS_ATTACH:
      SpawnCalc();
      break;
   default:
   case DLL_PROCESS_DETACH:
   case DLL_THREAD_ATTACH:
   case DLL_THREAD_DETACH:
      break;
   }
   return TRUE;
}