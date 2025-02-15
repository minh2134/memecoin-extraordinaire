/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'mystery': {
          dark: '#1E1E1E',
          accent: '#8A2BE2',
          highlight: '#FF6B6B',
        },
      },
      fontFamily: {
        'heading': ['DM Sans', 'sans-serif'],
        'body': ['Space Grotesk', 'sans-serif'],
      },
      keyframes: {
        float: {
          '0%, 100%': { transform: 'translateY(0)' },
          '50%': { transform: 'translateY(-20px)' },
        }
      },
      animation: {
        'float': 'float 6s ease-in-out infinite',
        'float-delayed': 'float 6s ease-in-out 3s infinite',
        'pulse': 'pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite',
      }
    },
  },
  plugins: [
    function({ addUtilities }) {
      const newUtilities = {
        '.border-gradient-to-r': {
          'border-image': 'linear-gradient(45deg, #E5E5E5, #8A2BE2) 1',
        },
      }
      addUtilities(newUtilities)
    }
  ],
} 