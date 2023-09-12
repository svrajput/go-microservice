
# Step1  : 
    ## Install mysql on local environment. 

# Step 2 :
   ## Login mysql server 

    command : use mysql;
    command : ALTER USER 'root'@'localhost' IDENTIFIED BY 'root.123';
    command : flush privileges;
    command : exit;

 # Step 3 :
    Connect : mysql -u root -p  

    
# Create database; 
CREATE DATABASE test_MS;
 
 
 ## Create Table 

CREATE TABLE `customers` (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `date_of_birth` date NOT NULL,
  `city` varchar(100) NOT NULL,
  `zipcode` varchar(10) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

# Insert Records
INSERT INTO `customers` VALUES
	(2000,'Raj','1978-12-15','Delhi','110075',1),
	(2001,'Arian','1988-05-21','Newburgh, NY','12550',1),
	(2002,'Hadley','1988-04-30','Englewood, NJ','07631',1),
	(2003,'Ben','1988-01-04','Manchester, NH','03102',0),
	(2004,'Nina','1988-05-14','Clarkston, MI','48348',1),
	(2005,'Rajesh','1988-11-08','Hyattsville, MD','20782',0);

