rm -rf dist
rm -rf ooppwwqq0_minify.egg-info
go get -u all
python -m build --sdist
twine upload dist/*
