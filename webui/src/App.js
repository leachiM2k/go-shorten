import React from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import {Alert, Col, Layout, Row, Typography} from 'antd';
import './App.less';
import CustomContent from './global/CustomContent';
import GlobalProvider from './context/GlobalProvider';
import Logout from './components/Logout';

const { Header, Content, Footer } = Layout;

const App = () => (
    <Router basename={process.env.PUBLIC_URL || '/ui'}>
        <GlobalProvider>
            <Layout>
                <Header className="header">
                    <Row>
                        <Col flex="1"><Typography.Title style={{ color: 'white' }}>Go Shorten!</Typography.Title></Col>
                        <Col>
                        </Col>
                    </Row>
                </Header>
                <Content className="site-layout-background">
                    <Alert.ErrorBoundary>
                        <CustomContent/>
                    </Alert.ErrorBoundary>
                </Content>
                <Footer>
                    <Row wrap={true} justify="space-between">
                        <Col>
                            <Logout/>
                        </Col>
                        <Col>
                            &copy;2021 Created by Michael Rotmanov
                        </Col>
                    </Row>
                </Footer>
            </Layout>
        </GlobalProvider>
    </Router>
);

export default App;
