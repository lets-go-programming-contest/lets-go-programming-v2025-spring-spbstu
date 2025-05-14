echo "By default:"
first_command="go build -o main"
echo $first_command

eval $first_command
echo
./main
echo

echo "With -ldflags:"
second_command="go build -o main -ldflags=\"-X 'main.MyFavouritePockemon=Squirtle'\""
echo $second_command

eval $second_command
echo
./main