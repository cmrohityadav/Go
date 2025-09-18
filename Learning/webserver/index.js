
const express = require('express')
const app = express()
const port = 8000

app.use(express.json()); 
app.use(express.urlencoded({extended: true}));

app.get('/', (req, res) => {
  res.status(200).send("Welcome to Server for Golang")
})

app.get('/get', (req, res) => {
    res.status(200).json({message: "Hello from cmrohityadav.in"})
  })


app.post('/post', (req, res) => {
    let myJson = req.body;      // your JSON
	
	res.status(200).send(myJson);
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})
app.get('/apijson',(req,res)=>{
  let arrayData = [
    { name: "Amit", class: 1, doy: 2015 },
    { name: "Sneha", class: 2, doy: 2013 },
    { name: "Rahul", class: 3, doy: 2017 },
    { name: "Priya", class: 4, doy: 2016 },
    { name: "Vikram", class: 6, doy: 2014 },
    { name: "Kiran", class: 2, doy: 2018 },
    { name: "Neha", class: 1, doy: 2019 },
    { name: "Sandeep", class: 8, doy: 2012 },
    { name: "Meera", class: 9, doy: 2011 },
    { name: "Arjun", class: 10, doy: 2010 }
  ];


return res.status(200).json({message:"success",data:arrayData})
})
  

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})