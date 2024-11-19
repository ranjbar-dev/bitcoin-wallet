echo "Enter commit message: "
set /p COMMENT=
echo "Enter version: (2.0.0)"
set /p VERSION=

git add . 

git commit -am "%COMMENT%"

git push

git tag "%VERSION%"

git push origin "%VERSION%"

go list -m github.com/ranjbar-dev/bitcoin-wallet@%VERSION%