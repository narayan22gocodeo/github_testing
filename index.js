const axios = require('axios');
const express = require('express');
const cors = require('cors');

const app = express();
app.use(cors());
app.use(express.json());

// Store upcoming listings in memory (in production, use a database)
let upcomingListings = [];

// Function to fetch upcoming listings from Binance Launchpad
async function fetchBinanceUpcomingListings() {
    try {
        const response = await axios.get('https://www.binance.com/bapi/composite/v1/public/cms/article/catalog/list/query?catalogId=48&pageNo=1&pageSize=10');
        return response.data.data.articles.map(article => ({
            exchange: 'Binance',
            token: article.title.split(' ')[0],
            date: new Date(article.releaseDate),
            description: article.title,
            link: `https://www.binance.com/en/support/announcement/${article.code}`
        }));
    } catch (error) {
        console.error('Error fetching api:', error);
        return [];
    }
}

// Function to fetch upcoming listings from multiple sources
async function updateUpcomingListings() {
    try {
        const binanceListings = await fetchBinanceUpcomingListings();
        upcomingListings = [...binanceListings];
        
        // Sort by date
        upcomingListings.sort((a, b) => a.date - b.date);
    } catch (error) {
        console.error('Error updating listings:', error);
    }
}

// API endpoints
app.get('/api/upcoming-listings', (req, res) => {
    res.json(upcomingListings);
});

// Update listings every 30 minutes
setInterval(updateUpcomingListings, 30 * 60 * 1000);

// Initial fetch
updateUpcomingListings();

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Crypto tracker server running on port ${PORT}`);
});

