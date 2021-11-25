<p>Mail Test</p>
<?php
$ret = mail('to@test.local', 'Test Mail', 'This is a test mail.', 'From:from@test.local');
if ($ret) {
    echo 'success';
} else {
    echo 'fail';
}