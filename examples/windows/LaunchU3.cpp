#include <windows.h>
#include <stdio.h>
#include <iostream>

int main(int argc, char** argv)
{
	printf("[+] Loading DLL...\n");
	CHAR* dllname = "dwmapi-winexec.dll";

	HMODULE hLibrary = LoadLibraryA(dllname);
	if (hLibrary == NULL)
	{
		printf("[-] LoadLibraryA has failed: %d\n", GetLastError());
		return 1;
	}

	printf("[+] Handle of the loaded DLL: 0x%p\n", hLibrary);

	FARPROC gpa = GetProcAddress(hLibrary, MAKEINTRESOURCEA(1));
	if (gpa)
	{
		printf("[+] Module Address: 0x%p\n", gpa);
	}
	else
	{
		printf("[-] Last Error: %d\n", GetLastError());
		return 1;
	}

	if (FreeLibrary(hLibrary))
	{
		printf("[+] Library has been unloaded successfuly!\n");
	}
	
	return 0;
}