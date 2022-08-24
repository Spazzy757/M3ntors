import logo from './logo.svg';
import TBar from './Toolbar'
import './App.css';

function App() {
  return (
    <div className="App">
      <TBar />
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
      </header>
    </div>
  );
}

export default App;
