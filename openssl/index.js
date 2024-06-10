const crypto = require('crypto');
const md5 = text => {
return crypto
    .createHash('md5')
    .update(text)
    .digest();
}

const decrypt = (encryptedBase64, secretKey) => {
    secretKey = md5(secretKey);
    secretKey = Buffer.concat([secretKey, secretKey.slice(0, 8)]); // properly expand 3DES key from 128 bit to 192 bit
    const decipher = crypto.createDecipheriv('des-ede3', secretKey, '');
    let decrypted = decipher.update(encryptedBase64, 'base64');
    decrypted += decipher.final();
    return decrypted;
};

console.log("Decrypted text:", decrypt('U2FsdGVkX1/4F20YCC6K/bksaZLd2xmgO7qGZCLA8TktonSKtJBxIPhTTpUNO9oNOm5Re3VjuuN7C9xCYiwlZ7+iMXHuGSEBZVHhlQnS+567tnn2G2I3anR920BxEBHuQ9vcsPUGzhVHG6IqMUnEzANXaLA76sjQGXKyDISKgy7ru0w14MK6ALAivDKW7McMHDARMH4/onm3FgYWODlvn4hjlIqV7hzsTESj+BPHW1adAxU4tdBgyNg9po3zMBaa/WoopkInUsnGacSjYetdA+CSH7HUxkwX4fHHnGPs1zGA3s1P5OiCuy/goF9sY1rv3me3u+VhdRMkp622xhzKrnfGyebVRxXuKZHejALUU8dNlzEJbEhFpsZLAGNKrKC6wWtGXEDLG2lrvD1D6g0BYcjzjzD/j6uZoGwWJEnZ88dz9OeRSQH/CP8XqjhQSvQ47F1Yz+fpdLzOG0LgMY9YcKutZJ0siP0D7rlNq+8bIa/QbBE960S7b+GtEbrdl4HZx0catDg3sgvVu6bKR1wr9/rqIJhiJvmnZeTNwbTbHZYntmWHxISC+MNHrFkrUgTcqZyjMoW1qClBT1+Nj8Taaty7jYIvgpoRMF3CeLCE5WtsRmhh7Jq4wiXWyL+pb4joUlDGFVB6MNRF1Yot9IJHUF5rsf5nebhC3VbssM4SAKxOEchnG7ccIOawabfL/g+NsHdDdMZQfC49fAKb/H82+6KKm2aXQ/889WAPvec62NPsPXYEHTQYYcYVNdRap+e1MfsdPrn1pqVH759el53PjlcGNhVZvityd3HE9xsN8ElIHoCHDFIQN3kbTSsWNuc98L0wltcq4kDBF0locA0y8Hk5UtToS+AZvFtCD/AAsrAUTODyuSYoqV3g2VxWAmOiOxVJmAr99+aRpY4ZY6H8QXiLyuM8ZOE4+dCrlgtY/h82sgLGg6L9QSu9A723l5Byxxdpyqp7fjZqOiZs1ktd5RDjQBa0zbRIvZqaIUOb2uM=', 'Фантомный друг'));