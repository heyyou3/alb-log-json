package fetch_alb_log

// NOTE: fetch alb log return string slice
//       Convert stored S3 logs from gz to txt
//       obligation
//       - fetch s3 object(directory)
//       - convert gz file to string slice
