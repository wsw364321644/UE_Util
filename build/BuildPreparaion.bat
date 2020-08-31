call setupgopath.bat
set buildpath=%~dp0
cd %~dp0..\bin
set binpath=%cd%
cd ..\src\github.com\wsw364321644\unrealenginetools\preparation

%binpath%\packr.exe
go build -o %binpath%\preparation.exe
%binpath%\packr.exe clean

cd %buildpath%