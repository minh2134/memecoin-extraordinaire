import React, { useState, useEffect } from 'react';
import { Line } from 'react-chartjs-2';
import { FaSearch } from 'react-icons/fa';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

const cryptoList = [
  { 
    id: 'bitcoin', 
    symbol: 'BTC', 
    name: 'Bitcoin', 
    nicknames: ['btc', 'bitcoin', 'digital gold', 'crypto king']
  },
  { 
    id: 'dogecoin', 
    symbol: 'DOGE', 
    name: 'Dogecoin', 
    nicknames: ['doge', 'dogecoin', 'much wow']
  },
  { 
    id: 'shiba-inu', 
    symbol: 'SHIB', 
    name: 'Shiba Inu', 
    nicknames: ['shib', 'shiba', 'doge killer']
  },
  { 
    id: 'bonk', 
    symbol: 'BONK', 
    name: 'Bonk', 
    nicknames: ['bonk', 'solana dog']
  },
  { 
    id: 'pepe', 
    symbol: 'PEPE', 
    name: 'Pepe', 
    nicknames: ['pepe', 'frog coin']
  },
];

const formatPrice = (price) => {
  if (price === 0) return '$0.00';
  if (price < 0.01) return `$${price.toFixed(8)}`;
  if (price < 1) return `$${price.toFixed(4)}`;
  return `$${price.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })}`;
};

const Market = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedCrypto, setSelectedCrypto] = useState('bitcoin');
  const [cryptoData, setCryptoData] = useState([]);
  const [chartData, setChartData] = useState(null);
  const [sortConfig, setSortConfig] = useState({
    key: null,
    direction: 'asc'
  });

  useEffect(() => {
    fetchCryptoData();
    fetchChartData(selectedCrypto);
    const interval = setInterval(fetchCryptoData, 60000); // Update every minute
    return () => clearInterval(interval);
  }, [selectedCrypto]);

  const fetchCryptoData = async () => {
    try {
      const response = await fetch(
        `https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=${cryptoList.map(c => c.id).join(',')}&order=market_cap_desc&sparkline=false&price_change_percentage=1h,24h`
      );
      const data = await response.json();
      setCryptoData(data);
    } catch (error) {
      console.error('Error fetching crypto data:', error);
    }
  };

  const fetchChartData = async (cryptoId) => {
    try {
      const response = await fetch(
        `https://api.coingecko.com/api/v3/coins/${cryptoId}/market_chart?vs_currency=usd&days=365`
      );
      const data = await response.json();
      
      const chartData = {
        labels: data.prices.map(price => new Date(price[0]).toLocaleDateString()),
        datasets: [
          {
            label: cryptoList.find(c => c.id === cryptoId)?.name || cryptoId,
            data: data.prices.map(price => price[1]),
            borderColor: '#8A2BE2',
            backgroundColor: 'rgba(138, 43, 226, 0.1)',
            fill: true,
          },
        ],
      };
      setChartData(chartData);
    } catch (error) {
      console.error('Error fetching chart data:', error);
    }
  };

  const handleSort = (key) => {
    let direction = 'asc';
    if (sortConfig.key === key && sortConfig.direction === 'asc') {
      direction = 'desc';
    }
    setSortConfig({ key, direction });
  };

  const getSortedData = (data) => {
    if (!sortConfig.key) return data;

    return [...data].sort((a, b) => {
      if (sortConfig.key === 'name') {
        return sortConfig.direction === 'asc' 
          ? a.name.localeCompare(b.name)
          : b.name.localeCompare(a.name);
      }

      let aValue = a[sortConfig.key];
      let bValue = b[sortConfig.key];

      // Handle nested properties for percentage changes
      if (sortConfig.key === 'price_change_percentage_1h_in_currency' || 
          sortConfig.key === 'price_change_percentage_24h') {
        aValue = aValue || 0;
        bValue = bValue || 0;
      }

      if (aValue < bValue) return sortConfig.direction === 'asc' ? -1 : 1;
      if (aValue > bValue) return sortConfig.direction === 'asc' ? 1 : -1;
      return 0;
    });
  };

  const filteredCryptos = cryptoData.filter(crypto => {
    const searchLower = searchTerm.toLowerCase();
    const coinInfo = cryptoList.find(c => c.id === crypto.id);
    return coinInfo?.nicknames.some(nickname => nickname.includes(searchLower)) ||
           crypto.name.toLowerCase().includes(searchLower) ||
           crypto.symbol.toLowerCase().includes(searchLower);
  });

  const sortedAndFilteredCryptos = getSortedData(filteredCryptos);

  const SortHeader = ({ label, sortKey, align = 'left' }) => {
    const isSorted = sortConfig.key === sortKey;
    return (
      <th 
        className={`px-6 py-4 text-${align} text-white cursor-pointer hover:text-mystery-accent transition-colors`}
        onClick={() => handleSort(sortKey)}
      >
        <div className={`flex items-center ${align === 'right' ? 'justify-end' : ''} space-x-1`}>
          <span>{label}</span>
          {isSorted && (
            <span className="ml-1">
              {sortConfig.direction === 'asc' ? '↑' : '↓'}
            </span>
          )}
        </div>
      </th>
    );
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0B041A] to-mystery-dark p-8">
      <div className="container mx-auto">
        {/* Chart Section - Now with reduced height */}
        <div className="bg-[#1E1E1E] rounded-xl p-4 mb-6" style={{ height: '300px' }}>
          {chartData && (
            <Line
              data={chartData}
              options={{
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                  legend: {
                    position: 'top',
                    labels: { color: 'white' }
                  },
                },
                scales: {
                  x: { 
                    grid: { color: '#333' }, 
                    ticks: { 
                      color: 'white',
                      maxRotation: 0,
                      autoSkip: true,
                      maxTicksLimit: 12
                    } 
                  },
                  y: { 
                    grid: { color: '#333' }, 
                    ticks: { color: 'white' } 
                  }
                }
              }}
            />
          )}
        </div>

        {/* Search Bar */}
        <div className="relative mb-6">
          <FaSearch className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-400" />
          <input
            type="text"
            placeholder="Search by name or nickname..."
            className="w-full pl-12 pr-4 py-3 bg-[#1E1E1E] rounded-xl text-white focus:outline-none focus:ring-2 focus:ring-mystery-accent"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>

        {/* Crypto List */}
        <div className="bg-[#1E1E1E] rounded-xl overflow-hidden">
          <table className="w-full">
            <thead>
              <tr className="bg-mystery-accent/20">
                <th className="px-6 py-4 text-left text-white">#</th>
                <SortHeader label="Token" sortKey="name" />
                <SortHeader label="Price" sortKey="current_price" align="right" />
                <SortHeader label="1h %" sortKey="price_change_percentage_1h_in_currency" align="right" />
                <SortHeader label="24h %" sortKey="price_change_percentage_24h" align="right" />
                <SortHeader label="FDV" sortKey="fully_diluted_valuation" align="right" />
                <SortHeader label="Market Cap" sortKey="market_cap" align="right" />
              </tr>
            </thead>
            <tbody>
              {sortedAndFilteredCryptos.map((crypto, index) => (
                <tr
                  key={crypto.id}
                  className="border-t border-gray-800 hover:bg-mystery-accent/5 cursor-pointer"
                  onClick={() => setSelectedCrypto(crypto.id)}
                >
                  <td className="px-6 py-4 text-white">{index + 1}</td>
                  <td className="px-6 py-4">
                    <div className="flex items-center">
                      <img src={crypto.image} alt={crypto.name} className="w-8 h-8 mr-3" />
                      <span className="text-white">{crypto.name}</span>
                      <span className="text-gray-400 ml-2">{crypto.symbol.toUpperCase()}</span>
                    </div>
                  </td>
                  <td className="px-6 py-4 text-right text-white">
                    {formatPrice(crypto.current_price)}
                  </td>
                  <td className={`px-6 py-4 text-right ${crypto.price_change_percentage_1h_in_currency > 0 ? 'text-green-500' : 'text-red-500'}`}>
                    {crypto.price_change_percentage_1h_in_currency?.toFixed(2) || '0.00'}%
                  </td>
                  <td className={`px-6 py-4 text-right ${crypto.price_change_percentage_24h > 0 ? 'text-green-500' : 'text-red-500'}`}>
                    {crypto.price_change_percentage_24h?.toFixed(2) || '0.00'}%
                  </td>
                  <td className="px-6 py-4 text-right text-white">
                    ${crypto.fully_diluted_valuation?.toLocaleString() || 'N/A'}
                  </td>
                  <td className="px-6 py-4 text-right text-white">
                    ${crypto.market_cap?.toLocaleString() || 'N/A'}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default Market;
export { cryptoList };
