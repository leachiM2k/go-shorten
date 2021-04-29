import React from "react";
import {Card, List} from 'antd';

const data = [
    {
        title: "Start time",
        content: "Enter the start time from which the links should become valid."
    },
    {
        title: "Expiration time",
        content: "Do you want the links to expire at a certain time? You can set it freely!"
    },
    {
        title: "Maximal visits limit",
        content: "After a defined number of requests, the links become invalid."
    },
    {
        title: "Detailed statistics",
        content: "`Go Shorten!` shows you the IP, the browser version and the time of access to your shortened links."
    },
];

const Features = props => {
    return (
        <List
            grid={{ gutter: 16, column: 2 }}
            dataSource={data}
            renderItem={item => (
                <List.Item>
                    <Card headStyle={{background:'#006d77',color:'#edf6f9'}} bodyStyle={{background:'#def9fc'}} title={item.title}>{item.content}</Card>
                </List.Item>
            )}
        />
    );
};

export default Features;
