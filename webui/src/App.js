import React from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import {Col, Layout, Row, Typography, Alert} from 'antd';
import './App.css';
import CustomContent from './global/CustomContent';
import GlobalProvider from './context/GlobalProvider';
import CustomRightHeader from './global/CustomRightHeader';

const { Header, Content } = Layout;

const App = () => (
    <Router>
        <GlobalProvider>
            <Layout>
                <Header className="header">
                    <Row>
                        <Col flex={1}><Typography.Title style={{ color: 'white' }}>Go Shorten!</Typography.Title></Col>
                        <Col>
                            <CustomRightHeader/>
                        </Col>
                    </Row>
                </Header>
                <Content className="site-layout-background">
                    <Alert.ErrorBoundary>
                        <CustomContent/>
                    </Alert.ErrorBoundary>
                </Content>
            </Layout>
        </GlobalProvider>
    </Router>
);

export default App;
