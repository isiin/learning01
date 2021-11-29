<?php

$mysqli = new mysqli('db', 'dbuser', 'dbpass', 'testdb');
if ($mysqli->connect_error) {
    echo $mysqli->connect_error;
} else {
    echo "Success";
    $mysqli->close();
}
