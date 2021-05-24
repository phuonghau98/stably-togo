import React, { useEffect, useState } from 'react';
import { Alert, Button, Col, Form, Input, InputNumber, Row } from 'antd';
import './App.css';

const FIND_PRIME_NUMBER_ENDPOINT =
  process.env.NODE_ENV === 'production'
    ? '/api/v2/prime/findnearest'
    : `http://localhost:8080/api/v2/prime/findnearest`;

function App() {
  const [errorMsg, setErrorMsg] = useState('');
  const [responsePrimeNumber, setResponseNumber] = useState<string | null>();
  const [requestedNumber, setRequestedNumber] = useState<string | null>();
  const inputRef = React.useRef<Input | null>(null);
  const [isFetchingResult, setFetchingResult] = useState(false);
  function handleFormSubmmited(values: { num: string }) {
    // async function fetch(): Promise<any> {
    //   return new Promise((resolve, reject) => {
    //     setTimeout(() => {
    //       const randomState = Math.floor(Math.random() * 10) + 1;
    //       console.log(randomState);
    //       if (randomState < 5) resolve({ num: -1 })
    //       if (randomState % 2 === 0) resolve({ num: randomState });
    //       else reject({ error: 'Something went wrong' });
    //     }, 500);
    //   });
    // }

    // restart state
    setResponseNumber(null);
    setRequestedNumber(values.num);
    setFetchingResult(true);
    setErrorMsg('');

    // Fetch response
    fetch(FIND_PRIME_NUMBER_ENDPOINT, {
      body: JSON.stringify({ num: values.num }),
      method: 'POST',
      
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.error) {
          setErrorMsg(data.error);
        }
        if (data && data.data) {
          setResponseNumber(data.data.num);
          setErrorMsg('');
        }
      })
      .catch((err) => {
        setErrorMsg(err.error ? err.error : err.toString());
      })
      .finally(() => {
        setFetchingResult(false);
      });
  }

  // Focus input again after receiving response
  useEffect(() => {
    if (!isFetchingResult && inputRef.current) {
      inputRef.current.focus()
    }
  }, [isFetchingResult])

  // Focus input after mounted
  useEffect(() => {
    if(inputRef.current) {
      inputRef.current.focus()
    }
  }, [])
  return (
    <div className="App">
      <Row justify="center" align="middle" style={{ height: '100%' }}>
        <Col style={{ maxWidth: 600 }}>
          <Form layout="vertical" onFinish={handleFormSubmmited}>
            <Form.Item
              label="Enter a number that you want to find the lower nearest prime"
              name="num"
              hasFeedback
              initialValue={"2"}
              rules={[
                {
                  required: true,
                  pattern: /^[0-9]{1,127}$/,
                  message: "Invalid number, input number has length from 1 to 127 digits"
                }
              ]}
            >
              <Input
                ref={inputRef}
                disabled={isFetchingResult}
                size="large"
                style={{ width: '100%' }}
              />
            </Form.Item>
            <Button
              loading={isFetchingResult}
              type="primary"
              htmlType="submit"
              size="large"
              style={{ width: '100%' }}
            >
              {isFetchingResult ? 'Figuring it out...' : 'Find'}
            </Button>
          </Form>
          <div style={{ minHeight: 40, marginTop: 20 }}>
            {errorMsg ? (
              <Alert message={<>{errorMsg}</>} type="error" />
            ) : (
              requestedNumber &&
              (responsePrimeNumber === '-1' ? (
                <Alert message="No prime number found" type="info" />
              ) : (
                responsePrimeNumber && (
                  <Alert
                    message={
                      <>
                        The highest prime number lower than {requestedNumber}{' '}
                        is: <b>{responsePrimeNumber}</b>
                      </>
                    }
                    type="success"
                  />
                )
              ))
            )}
          </div>
        </Col>
      </Row>
    </div>
  );
}

export default App;
