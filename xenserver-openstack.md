<table>
    <tbody>
        <tr>
			<td>id</td>
            <td>action</td>
            <td>support</td>
        </tr>
        <tr>
            <td>1</td>
            <td>suspend</td>
            <td>ok</td>
        </tr>
        <tr>
            <td>2</td>
            <td>resume</td>
            <td>ok</td>
        </tr>
        <tr>
            <td>3</td>
            <td>view log</td>
            <td>ok</td>
        </tr>
        <tr>
            <td>4</td>
            <td>console</td>
            <td>failed</td>
        </tr>
        <tr>
            <td>5</td>
            <td>lock</td>
            <td>*</td>
        </tr>
        <tr>
            <td>6</td>
            <td>unlock</td>
            <td>*</td>
        </tr>
        <tr>
            <td>7</td>
            <td>soft reboot</td>
            <td>ok</td>
        </tr>
        <tr>
            <td>8</td>
            <td>hard reboot</td>
            <td>ok</td>
        </tr>
        <tr>
            <td>9</td>
            <td>shut off</td>
            <td>ok</td>
        </tr>
        <tr>
            <td>10</td>
            <td>terminate</td>
            <td>ok</td>
        </tr>
        <tr>
            <td>11</td>
            <td>start</td>
            <td>ok</td>
        </tr>
        <tr>
            <td>12</td>
            <td>migration</td>
            <td>failed</td>
        </tr>
        <tr>
            <td>13</td>
            <td>associate floating ip</td>
            <td>*</td>
        </tr>
        <tr>
            <td>14</td>
            <td>edit security group</td>
            <td>*</td>
        </tr>
        <tr>
            <td>15</td>
            <td>create snapshot</td>
            <td>ok</td>
        </tr>
        <tr>
            <td>16</td>
            <td>resize </td>
            <td>ok</td>
        </tr>
        <tr>
            <td>17</td>
            <td>rebuild</td>
            <td>ok</td>
        </tr>
    </tbody>
</table>

未解决的问题：

+ horizon上console窗口显示失败
* live-migration失败
* 对于ovs的neutron，对于network和sub-network不会创建对应qr-xxx之类的.
	* nova.conf里在xenserver.ovs_integration_bridge配置项，该配置是xenserver上一个network，对应ovs上的一个桥。
	* 之后创建的虚拟机的网卡都直接创建在这个桥上
* 由于上述网络问题，导致安全组是不可用


