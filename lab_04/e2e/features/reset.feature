Feature: Reset password with 2FA

Scenario: User reset password with 2FA
  When User send "POST" request to "/reset"
  Then the response on /reset code should be 200
  And the response on /reset should match json:
      """
      {
        "success": "Now should enter 2fa code to reset password"
      }
      """
  And user send "POST" request to "/verifyResetPassword"
  Then the response on /verifyResetPassword code should be 200
  And the response on /verifyResetPassword should match json:
      """
      {
        "success": "successful changing pass"
      }
      """