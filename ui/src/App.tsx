import React from 'react';
import logo from './logo.svg';
import { Form, Input } from 'antd';
import './App.css'

function App() {
  return (
    <div className="App">
      <Form>
        <Form.Item label="Enter a number that you want to find the lower nearest prime">
          <Input />
        </Form.Item>
      </Form>
    </div>
  );
}

export default App;
