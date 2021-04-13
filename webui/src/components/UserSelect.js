import {message, Select, Spin} from 'antd';
import client from '../lib/client-fetch';
import React, {useEffect, useState} from 'react';

const { Option } = Select;

const UserSelect = (props) => {
    let [user, setUser] = useState([]);
    let [loading, setLoading] = useState(true);

    const makeDataRequest = async () => {
        setLoading(true);
        try {
            const result = await client.get('/api/teams/members');
            if (result.data) {
                const data = result.data.map(user => ({
                    text: user.name + (user.teamid ? ' (Team: ' + user.teamid + ')' : ''),
                    value: user.pnumber,
                })).sort((a,b) => a.text.localeCompare(b.text));
                setUser(data);
            } else {
                setUser([]);
            }
        } catch (error) {
            message.error('Request failed: ' + error.message);
        }
        setLoading(false);
    }

    useEffect(() => { makeDataRequest() }, []);

    return (
        <Select
            mode="multiple"
            labelInValue
            value={props.selected || []}
            placeholder="Select users"
            notFoundContent={loading ? <Spin size="small"/> : null}
            filterOption={(input, option) =>
                option.children.toLowerCase().indexOf(input.toLowerCase()) >= 0 || option.value.includes(input)
            }
            onSelect={props.onSelect}
            style={{ width: '100%' }}
        >
            {user.map(d => (
                <Option key={d.value} value={d.value}>{d.text}</Option>
            ))}
        </Select>
    );
}

export default UserSelect;
