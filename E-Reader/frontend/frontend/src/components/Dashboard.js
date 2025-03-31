import React, { useState } from "react";
import { AppBar, Toolbar, Typography, Container, TextField, Grid, Card, CardMedia, CardContent } from "@mui/material";


const Dashboard = () => {
    const [searchTerm, setSearchTerm] = useState("");
    const books = [
        { id: 1, title: "A Court of Thorns and Roses", author: "Sarah J. Maas", image: "https://via.placeholder.com/150" },
        { id: 2, title: "Everything I Know About Love: A Memoir", author: "Dolly Alderton", image: "https://via.placeholder.com/150" },
        { id: 3, title: "Funny Story", author: "Emily Henry", image: "https://via.placeholder.com/150" },
        { id: 4, title: "The Heaven & Earth Grocery Store: A Novel", author: "James McBride", image: "https://via.placeholder.com/150" },
        { id: 5, title: "Crying in H Mart", author: "Michelle Zauner", image: "https://via.placeholder.com/150" },
        { id: 6, title: "James: A Novel", author: "Percival Everett", image: "https://via.placeholder.com/150" },
        { id: 7, title: "The Midnight Library: A GMA Book Club Pick: A Novel", author: "Matt Haig", image: "https://via.placeholder.com/150" },
        { id: 8, title: "Dog Man: The Scarlet Shedder", author: "Dav Pilkey", image: "https://via.placeholder.com/150" },
    ];

    const filteredBooks = books.filter((book) =>
        book.title.toLowerCase().includes(searchTerm.toLowerCase())
    );

    return (
        <div className="dashboard-container">
            {/* Header */}
            <AppBar position="static">
                <Toolbar>
                   <Typography variant="h6">E-Reader</Typography>
               </Toolbar>
            </AppBar>

            {/* Main Content */}
            <Container>
                <Typography variant="h4" gutterBottom>My Library</Typography>
                <TextField
                    fullWidth
                    label="Search for a book..."
                    variant="outlined"
                    value={searchTerm}
                    onChange={(e) => setSearchTerm(e.target.value)}
                    margin="normal"
                />
                <Grid container spacing={3}>
                    {filteredBooks.length > 0 ? (
                        filteredBooks.map((book) => (
                            <Grid item xs={12} sm={6} md={3} key={book.id}>
                                <Card>
                                    <CardMedia component="img" height="200" image={book.image} alt="Book Cover" />
                                    <CardContent>
                                        <Typography variant="h6">{book.title}</Typography>
                                        <Typography variant="body2" color="textSecondary">Author: {book.author}</Typography>
                                    </CardContent>
                                </Card>
                            </Grid>
                        ))
                    ) : (
                        <Typography variant="h6" color="error">No books found.</Typography>
                    )}
                </Grid>
            </Container>

            {/* Footer */}
            <AppBar position="static" color="primary" className="footer">
                <Toolbar>
                    <Typography variant="body2" color="inherit">&copy; 2025 E-Reader. All rights reserved.</Typography>
                </Toolbar>
            </AppBar>
        </div>
    );
};

export default Dashboard;
